import { calData, spdxData } from "licenses";

export default function Page(): JSX.Element {
  return (
    <div className="page">
      {spdxData.map((i) => (
        <pre key={i.licenseId}>
          {i.name} - {i.licenseId}
          <p>
            {calData.find((j) => j["spdx-id"] === i.licenseId)?.description}
          </p>
        </pre>
      ))}
    </div>
  );
}
