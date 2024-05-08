import { ClientResponseError } from '$lib/api/ClientResponseError';
import { toast } from 'svelte-sonner';

export const toastError = async (error: any) => {
	let errorString;
	if (typeof error == 'string') {
		errorString = error;
	} else if (error instanceof ClientResponseError) {
		if (error.status == 401) return;
		else if (error.status == 403) {
			if (error.message == 'invalid claim') {
				errorString = 'Insufficient credentials';
			} else {
				errorString = error.message;
			}
		} else errorString = error.message;
	} else if (error.isSuperTokensGeneralError === true) {
		errorString = error.message;
	} else if (error?.status >= 400 && error?.status < 500) {
		errorString = await error.json();
	} else errorString = 'Something went wrong. Please try again.';
	toast.error(errorString);
};
export const toastSuccess = (message: string) => {
	toast.success(message);
};
