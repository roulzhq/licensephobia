import {
  LicenseRulePermission,
  LicenseRuleCondition,
  LicenseRuleLimitation,
} from "types";

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

export interface CALLicenseDefinition {
  title: string;
  "spdx-id": string;
  description: string;
  how: string;
  using: Record<string, string> | null;
  permissions: LicenseRulePermission[];
  conditions: LicenseRuleCondition[];
  limitations: LicenseRuleLimitation[];
}
