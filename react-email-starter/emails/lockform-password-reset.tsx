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

interface LockformPasswordResetEmailProps {
  passwordResetLink?: string;
}

export const LockformPasswordResetEmail = ({
  passwordResetLink,
}: LockformPasswordResetEmailProps) => {
  const previewText = `Reset your password.`;

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
              Someone recently requested a password change for your Lockform
              account. If this was you, you can set a new password here:
            </Text>
            <Section className="text-center mt-[32px] mb-[32px]">
              <Button
                className="bg-[#000000] rounded text-white text-[12px] font-semibold no-underline text-center px-5 py-3"
                href={passwordResetLink}
              >
                Reset password
              </Button>
            </Section>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              If you don't want to change your password or didn't request this,
              just ignore and delete this message.
            </Text>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              To keep your account secure, please don't forward this email to
              anyone.
            </Text>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              The Lockform team
            </Text>
          </Container>
        </Body>
      </Tailwind>
    </Html>
  );
};

LockformPasswordResetEmail.PreviewProps = {
  email: "cleophas.barwareng@gmail.com",
  passwordResetLink: "https://app.lockform.io",
} as LockformPasswordResetEmailProps;

export default LockformPasswordResetEmail;
