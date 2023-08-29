export type PackageManagerId = "npm"

export interface DependencyFileScanResult {
  packageManager: PackageManagerId;
  results: PackageScanResult[];
}

export interface PackageScanResult {
  name: string;
  version: string;
  licenseId: string;
  websiteUrl: string;
}