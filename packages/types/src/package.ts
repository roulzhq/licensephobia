export type PackageManagerId = "npm";

export interface DependencyFileScanResult {
  packageManager: PackageManagerId;
  results: PackageScanResult[];
}

export interface PackageScanResult {
  id: string;
  source: string;
  version: string;
}

export interface PackageCheckResult {
  id: string;
  description: string;
  name: string;
  licenseId: string;
  homepage: string;
}
