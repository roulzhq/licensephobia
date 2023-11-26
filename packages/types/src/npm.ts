export interface PackageJson {
  dependencies: Record<string, string>;
  devDependencies: Record<string, string>;
  peerDependencies: Record<string, string>;
  description: string;
  license: string;
  maintainers: NpmRegPackageMaintainer[];
  homepage: string;
  [other: string]: unknown;
}

export interface NpmRegPackage extends PackageJson {
  _id: string;
  name: string;
}

export interface NpmRegPackageVersion extends PackageJson {
  name: string;
  version: string;
  author: string;
}

export interface NpmRegPackageMaintainer {
  name: string;
  email: string;
}
