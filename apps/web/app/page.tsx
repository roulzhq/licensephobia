import { getLicenseDefinitionById } from "core";

export default function Page(): JSX.Element {
  return (
    <div className="page">
      <pre>{JSON.stringify(getLicenseDefinitionById("0BSD"), null, 2)}</pre>
    </div>
  );
}
