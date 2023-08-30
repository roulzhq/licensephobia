export interface PackageJsonInput {
  dependencies: Record<string, string>;
  devDependencies: Record<string, string>;
  peerDependencies: Record<string, string>;
  [other: string]: unknown;
}

export interface NpmRegPackageResponse {
  name: "turbo",
  
}