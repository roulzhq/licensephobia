export type PackageManagerId = "npm";

export interface DependencyFileScanResult {
  packageManager: PackageManagerId;
  results: PackageScanResult[];
}

export interface PackageScanResult {
  name: string;
  source: string;
  version: string;
  licenseId: string | null;
  websiteUrl: string | null;
}
