import { LicenseDefinition } from "@licensephobia/types";
import { spdxData, calData } from "@licensephobia/licenses";

export function getLicenseDefinitionById(
  licenseId: string
): LicenseDefinition | undefined {
  let license: LicenseDefinition = {
    id: licenseId,
    calFound: false,
    spdxFound: false,
  };

  const spdxLicense = spdxData.find(
    (i) => i.licenseId.toLowerCase() === licenseId.toLowerCase()
  );
  const calLicense = calData.find(
    (i) => i["spdx-id"].toLowerCase() === licenseId.toLowerCase()
  );

  if (!spdxLicense && !calLicense) return license;

  if (spdxLicense) {
    license.spdxFound = true;
    license.name = spdxLicense.name;
    license.spdxDetailsUrl = spdxLicense.detailsUrl;
    license.spdxReferenceUrl = spdxLicense.reference;
    license.seeAlsoUrls = spdxLicense.seeAlso;
    license.isDeprecatedLicenseId = spdxLicense.isDeprecatedLicenseId;
  }

  if (calLicense && !spdxLicense) {
    license.name = calLicense.title;
  }

  if (calLicense) {
    license.calFound = true;
    license.description = calLicense.description;
    license.how = calLicense.how;

    license.permissions = calLicense.permissions;
    license.conditions = calLicense.conditions;
    license.limitations = calLicense.limitations;
  }

  return license;
}

export function getAllLicenses(): LicenseDefinition[] {
  return spdxData
    .map((i) => getLicenseDefinitionById(i.licenseId))
    .filter((i) => i !== undefined) as LicenseDefinition[];
}
