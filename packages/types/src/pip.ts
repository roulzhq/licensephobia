export interface PipfileInput {
  "dev-packages": Record<string, string>;
  packages: Record<string, string>;
  [other: string]: unknown;
}
