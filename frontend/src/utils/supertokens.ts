import SuperTokens from 'supertokens-web-js';
import EmailVerification, {
	sendVerificationEmail,
	verifyEmail
} from 'supertokens-web-js/recipe/emailverification';

import Session from 'supertokens-web-js/recipe/session';
import ThirdPartyEmailPassword, {
	emailPasswordSignIn,
	emailPasswordSignUp,
	getAuthorisationURLWithQueryParamsAndSetState,
	sendPasswordResetEmail,
	submitNewPassword,
	thirdPartySignInAndUp
} from 'supertokens-web-js/recipe/thirdpartyemailpassword';
import {
	VITE_SUPERTOKENS_COOKIE_DOMAIN,
	VITE_API_BASE_URL,
	VITE_APP_BASE_URL,
	VITE_SUPERTOKENS_APP_NAME
} from '$lib/env';
import { goto, invalidateAll } from '$app/navigation';
import { deleteTeamCookie } from '$utils';
import { toastError, toastSuccess } from './toasts';
// recei
export const supertokensInit = () => {
	SuperTokens.init({
		appInfo: {
			apiDomain: VITE_API_BASE_URL as string,
			apiBasePath: '/auth',
			appName: VITE_SUPERTOKENS_APP_NAME as string
		},
		recipeList: [
			EmailVerification.init(),
			Session.init({
				autoAddCredentials: true
				// sessionTokenBackendDomain: (VITE_SUPERTOKENS_COOKIE_DOMAIN as string) ?? undefined,
				// sessionTokenFrontendDomain: (VITE_SUPERTOKENS_COOKIE_DOMAIN as string) ?? undefined,
			}),
			ThirdPartyEmailPassword.init()
		]
	});
};

export const signupWithEmailAndPassword = async (email: string, password: string) => {
	let emailErrors: string[] = [];
	let passwordErrors: string[] = [];

	try {
		const response = await emailPasswordSignUp({
			formFields: [
				{
					id: 'email',
					value: email
				},
				{
					id: 'password',
					value: password
				}
			]
		});

		if (response.status === 'FIELD_ERROR') {
			response.formFields?.forEach((field) => {
				if (field.id == 'email') {
					emailErrors = emailErrors.concat(field.error);
				}
				if (field.id == 'password') {
					passwordErrors = passwordErrors.concat(field.error);
				}
			});
			return { emailErrors, passwordErrors };
		} else {
			await invalidateAll();
			await sendEmailVerificationLink();
		}
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}

	return { emailErrors, passwordErrors };
};

export const signinWithEmailAndPassword = async (email: string, password: string) => {
	try {
		const response = await emailPasswordSignIn({
			formFields: [
				{
					id: 'email',
					value: email
				},
				{
					id: 'password',
					value: password
				}
			]
		});

		let emailErrors: string[] = [];
		let passwordErrors: string[] = [];
		if (response.status === 'FIELD_ERROR') {
			response.formFields?.forEach((field) => {
				if (field.id == 'email') {
					emailErrors = emailErrors.concat(field.error);
				}
				if (field.id == 'password') {
					passwordErrors = passwordErrors.concat(field.error);
				}
			});
		} else if (response.status === 'WRONG_CREDENTIALS_ERROR') {
			// TODO display error
			passwordErrors = passwordErrors.concat('Email password combination is incorrect.');
		} else {
			goto('/', { invalidateAll: true });
		}
		return { emailErrors, passwordErrors };
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}
};

export const oauthLogin = async (thirdPartyId: 'google' | 'github') => {
	try {
		const authUrl = await getAuthorisationURLWithQueryParamsAndSetState({
			thirdPartyId,
			frontendRedirectURI: `${VITE_APP_BASE_URL}/oauth-callback/${thirdPartyId}`
		});
		window.location.href = authUrl;
	} catch (err: any) {
		toastError(err);
	}
};

export const handleOauthCallback = async () => {
	try {
		const response = await thirdPartySignInAndUp();

		if (response.status === 'OK') {
			if (response.createdNewRecipeUser) {
				// Add user to DB
				goto('/verify-email/success', { invalidateAll: true });
			} else {
				// Go to onboarding if not onboarded, otherwise go to home page
				goto('/', { invalidateAll: true });
			}
		} else {
			// SuperTokens requires that the third party provider
			// gives an email for the user. If that's not the case, sign up / in
			// will fail.

			// As a hack to solve this, you can override the backend functions to create a fake email for the user.
			goto('/signin', { invalidateAll: true });
		}
	} catch (err: any) {
		toastError(err);
		goto('/signin', { invalidateAll: true });
	}
};

export const sendEmailVerificationLink = async () => {
	try {
		const response = await sendVerificationEmail();
		if (response.status === 'EMAIL_ALREADY_VERIFIED_ERROR') {
			goto('/', { invalidateAll: true });
		} else {
			goto('/verify-email', { invalidateAll: true });
		}
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}
};

export const consumeVerificationCode = async () => {
	try {
		const response = await verifyEmail();
		if (response.status === 'EMAIL_VERIFICATION_INVALID_TOKEN_ERROR') {
			goto('/verify-email/failed', { invalidateAll: true });
		} else {
			goto('/verify-email/success');
		}
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}
};

export const sendResetPasswordLink = async (email: string) => {
	try {
		const response = await sendPasswordResetEmail({
			formFields: [
				{
					id: 'email',
					value: email
				}
			]
		});
		let emailErrors: string[] = [];
		if (response.status === 'FIELD_ERROR') {
			response.formFields?.forEach((field) => {
				if (field.id == 'email') {
					emailErrors = emailErrors.concat(field.error);
				}
			});
			return emailErrors;
		} else {
			goto('/reset-password/link-sent');
		}
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}
};

export const newPasswordEntered = async (newPassword: string) => {
	try {
		const response = await submitNewPassword({
			formFields: [
				{
					id: 'password',
					value: newPassword
				}
			]
		});

		let passwordErrors: string[] = [];
		if (response.status === 'FIELD_ERROR') {
			response.formFields?.forEach((field) => {
				if (field.id == 'password') {
					passwordErrors = passwordErrors.concat(field.error);
				}
			});
			return passwordErrors;
		} else if (response.status === 'RESET_PASSWORD_INVALID_TOKEN_ERROR') {
			toastError('The password reset token is ivalid or expired.');
			goto('/signin', { invalidateAll: true });
		} else {
			toastSuccess('Password updated successfully');
			goto('/', { invalidateAll: true });
		}
	} catch (err: any) {
		toastError(err);
		if (err?.status >= 400 && err?.status < 500) goto('/signin', { invalidateAll: true });
	}
};

export const logout = async () => {
	await Session.signOut();
	deleteTeamCookie();
	goto('/signin', { invalidateAll: true });
};
