import {
  Body,
  Button,
  Container,
  Column,
  Head,
  Heading,
  Hr,
  Html,
  Img,
  Link,
  Preview,
  Row,
  Section,
  Text,
  Tailwind,
} from "@react-email/components";
import * as React from "react";

interface LockformInviteUserEmailProps {
  firstName?: string;
  email: string;
  userImage?: string;
  invitedByFirstName?: string;
  invitedByEmail?: string;
  teamName?: string;
  teamImage?: string;
  inviteLink?: string;
}

const baseUrl = process.env.VERCEL_URL
  ? `https://${process.env.VERCEL_URL}`
  : "";

export const LockformInviteUserEmail = ({
  firstName,
  email,
  invitedByFirstName,
  invitedByEmail,
  teamName,
  inviteLink,
}: LockformInviteUserEmailProps) => {
  const previewText = `Join ${invitedByFirstName} on Lockform`;

  return (
    <Html>
      <Head />
      <Preview>{previewText}</Preview>
      <Tailwind>
        <Body className="bg-white my-auto mx-auto font-sans px-2">
          <Container className="border border-solid border-[#eaeaea] rounded my-[40px] mx-auto p-[20px] max-w-[465px]">
            <Section className="mt-[32px]">
              <Img
                src={`${baseUrl}/static/vercel-logo.png`}
                width="40"
                height="37"
                alt="Lockform"
                className="my-0 mx-auto"
              />
            </Section>
            <Heading className="text-[#0A0A0A] text-[24px] font-normal text-center p-0 my-[30px] mx-0">
              Join <strong>{teamName}</strong> on <strong>Lockform</strong>
            </Heading>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              Hello {firstName},
            </Text>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              <strong>{invitedByFirstName}</strong> (
              <Link
                href={`mailto:${invitedByEmail}`}
                className="text-blue-600 no-underline"
              >
                {invitedByEmail}
              </Link>
              ) has invited you to the <strong>{teamName}</strong> team on{" "}
              <strong>Lockform</strong>.
            </Text>
            <Section className="text-center mt-[32px] mb-[32px]">
              <Button
                className="bg-[#000000] rounded text-white text-[12px] font-semibold no-underline text-center px-5 py-3"
                href={inviteLink}
              >
                Join the team
              </Button>
            </Section>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              or copy and paste this URL into your browser:{" "}
              <Link href={inviteLink} className="text-blue-600 no-underline">
                {inviteLink}
              </Link>
            </Text>
            <Hr className="border border-solid border-[#eaeaea] my-[26px] mx-0 w-full" />
            <Text className="text-[#666666] text-[12px] leading-[24px]">
              This invitation was intended for{" "}
              <Link
                href={`mailto:${email}`}
                className="text-blue-600 no-underline"
              >
                {invitedByEmail}
              </Link>
              . If you were not expecting this invitation, you can ignore this
              email. If you are concerned about your account's safety, please
              reply to this email to get in touch with us.
            </Text>
          </Container>
        </Body>
      </Tailwind>
    </Html>
  );
};

LockformInviteUserEmail.PreviewProps = {
  firstName: "Cleophas",
  email: "cleophas.barwareng@gmail.com",
  invitedByFirstName: "James",
  invitedByEmail: "james@poised.com",
  teamName: "Poised",
  inviteLink: "https://vercel.com/teams/invite/foo",
} as LockformInviteUserEmailProps;

export default LockformInviteUserEmail;
