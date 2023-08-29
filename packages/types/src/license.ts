
export type LicenseRulePermission =
  | "commercial-use"
  | "modifications"
  | "distribution"
  | "private-use"
  | "patent-use";

export type LicenseRuleCondition =
  | "include-copyright"
  | "include-copyright--source"
  | "document-changes"
  | "disclose-source"
  | "network-use-disclose"
  | "same-license"
  | "same-license--file"
  | "same-license--library";

export type LicenseRuleLimitation =
  | "trademark-use"
  | "liability"
  | "patent-use"
  | "warranty";

export interface LicenseDefinition {
  id: string;
  spdxFound: boolean;
  calFound: boolean;
  name?: string;
  description?: string;
  how?: string;
  spdxReferenceUrl?: string;
  spdxDetailsUrl?: string;
  seeAlsoUrls?: string[];
  isDeprecatedLicenseId?: boolean;
  permissions?: LicenseRulePermission[];
  conditions?: LicenseRuleCondition[];
  limitations?: LicenseRuleLimitation[];
}

export interface LicenseCheckRuleReason {
  sourceLicenseIds: string[];
}

export interface LicenseCheckResult {
  permissions: Record<LicenseRulePermission, LicenseCheckRuleReason>;
  conditions: Record<LicenseRuleCondition, LicenseCheckRuleReason>;
  limitations: Record<LicenseRuleLimitation, LicenseCheckRuleReason>;
}