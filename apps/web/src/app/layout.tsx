import "@fontsource/ibm-plex-mono";
import "$styles/base.scss";

import Header from "$components/header";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}): JSX.Element {
  return (
    <html lang="en">
      <body>
        <Header />
        <main>{children}</main>
      </body>
    </html>
  );
}
