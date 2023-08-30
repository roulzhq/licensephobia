"use client";

import { useSearchParams } from "next/navigation";

export default function Page(): JSX.Element {
  const searchParams = useSearchParams();
  const packages = searchParams.getAll("package");
  return (
    <div>
      <h1>scan</h1>
      <pre>{packages.join(", ")}</pre>
    </div>
  );
}
