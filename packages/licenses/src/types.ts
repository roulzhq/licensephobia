// SPDX Repository types

export interface SPDXLicenseDefinition {
  reference: string;
  isDeprecatedLicenseId: boolean;
  detailsUrl: string;
  referenceNumber: number;
  name: string;
  licenseId: string;
  seeAlso: Array<string>;
  isOsiApproved: boolean;
}

export interface SPDXResponse {
  licenseListVersion: string;
  licenses: SPDXLicenseDefinition[];
  releaseDate: string;
}

// Choosealicense.com repository types

export type CALPermission =
  | "commercial-use"
  | "modifications"
  | "distribution"
  | "private-use"
  | "patent-use";

export type CALCondition =
  | "include-copyright"
  | "include-copyright--source"
  | "document-changes"
  | "disclose-source"
  | "network-use-disclose"
  | "same-license"
  | "same-license--file"
  | "same-license--library";

export type CALLimitation =
  | "trademark-use"
  | "liability"
  | "patent-use"
  | "warranty";

export interface CALLicenseDefinition {
  title: string;
  "spdx-id": string;
  description: string;
  how: string;
  using: Record<string, string> | null;
  permissions: CALPermission[];
  conditions: CALCondition[];
  limitations: CALLimitation[];
}
