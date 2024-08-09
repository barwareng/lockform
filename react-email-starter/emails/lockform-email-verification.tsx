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

interface LockformVerifyEmailProps {
  verificationLink?: string;
}

export const LockformVerifyEmail = ({
  verificationLink,
}: LockformVerifyEmailProps) => {
  const previewText = `Please verify your email address.`;

  return (
    <Html>
      <Head />
      <Preview>{previewText}</Preview>
      <Tailwind>
        <Body className="bg-white my-auto mx-auto font-sans px-2">
          <Container className="border border-solid border-[#eaeaea] rounded my-[40px] mx-auto p-[20px] max-w-[465px]">
            <Section className="mt-[32px]">
              <Img
                src="https://poised.fra1.cdn.digitaloceanspaces.com/lockform-logo.png"
                width="40"
                height="37"
                alt="Lockform"
                className="my-0 mx-auto"
              />
            </Section>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px] mt-6">
              Hello,
            </Text>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              Welcome to Lockform, the platform that helps you be sure who you
              are communicating with. Please verify your email address to
              continue.
            </Text>
            <Section className="text-center mt-[32px] mb-[32px]">
              <Button
                className="bg-[#000000] rounded text-white text-[12px] font-semibold no-underline text-center px-5 py-3"
                href={verificationLink}
              >
                Verify Email
              </Button>
            </Section>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              or copy and paste this URL into your browser:{" "}
              <Link
                href={verificationLink}
                className="text-blue-600 no-underline"
              >
                {verificationLink}
              </Link>
            </Text>
            <Hr className="border border-solid border-[#eaeaea] my-[26px] mx-0 w-full" />
          </Container>
        </Body>
      </Tailwind>
    </Html>
  );
};

LockformVerifyEmail.PreviewProps = {
  email: "cleophas.barwareng@gmail.com",
  verificationLink: "https://app.lockform.io",
} as LockformVerifyEmailProps;

export default LockformVerifyEmail;
