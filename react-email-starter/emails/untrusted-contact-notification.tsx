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

interface UntrustedContactNotificationEmailProps {
  value: string;
  label: string;
  reasonForNotTrusting: string;
  addedByFirstName: string;
  addedByEmail: string;
  contactUrl: string;
}

export const UntrustedContactNotificationEmail = ({
  value,
  label,
  reasonForNotTrusting,
  addedByFirstName,
  addedByEmail,
  contactUrl,
}: UntrustedContactNotificationEmailProps) => {
  const previewText = `${value} is untrusted by your team.`;

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

            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              Hello,
            </Text>
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              <strong>{addedByFirstName}</strong> (
              <Link
                href={`mailto:${addedByEmail}`}
                className="text-blue-600 no-underline"
              >
                {addedByEmail}
              </Link>
              ) has untrusted the following contact on your team:
            </Text>
            <Section className="text-center">
              <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
                <strong>{value}</strong> [{label}]
              </Text>
            </Section>

            <Section className="text-center">
              <Button
                className="bg-[#000000] rounded text-white text-[12px] font-semibold no-underline text-center px-5 py-3"
                href={contactUrl}
              >
                View contact
              </Button>
              <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
                Please exercise if you communicate with them.
              </Text>
            </Section>

            <Hr className="border border-solid border-[#eaeaea] my-[26px] mx-0 w-full" />
            <Text className="text-[#0A0A0A] text-[14px] leading-[24px]">
              If you believe this was a mistake, please reach out to{" "}
              <strong>{addedByFirstName}</strong> (
              <Link
                href={`mailto:${addedByEmail}`}
                className="text-blue-600 no-underline"
              >
                {addedByEmail}
              </Link>
              )
            </Text>
          </Container>
        </Body>
      </Tailwind>
    </Html>
  );
};

UntrustedContactNotificationEmail.PreviewProps = {
  value: "email@example.com",
  label: "Budding Scammer",
  reasonForNotTrusting: "Out tryig to scam",
  addedByEmail: "user@acme.corp",
  addedByFirstName: "James",
  contactUrl: "http://127.0.0.1:5173/contacts/4",
} as UntrustedContactNotificationEmailProps;

export default UntrustedContactNotificationEmail;
