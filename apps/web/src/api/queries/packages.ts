import { PackageManagerId } from "@licensephobia/types";
import { checkPackages } from "@licensephobia/core";
import { useQuery } from "@tanstack/react-query";

export function useCheckPackages(
  packageManagerId: PackageManagerId,
  packages: [string, string][]
) {
  return useQuery({
    queryKey: ["CHECK_PACKAGES", packageManagerId, { packages }],
    queryFn: () => checkPackages(packageManagerId, packages),
    initialData: {},
  });
}
