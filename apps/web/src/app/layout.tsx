import "@fontsource/ibm-plex-mono";

import Header from "$components/header";

import styles from "./layout.module.scss";
import Providers from "./providers";

import "$styles/base.scss";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}): JSX.Element {
  return (
    <html lang="en">
      <body>
        <Providers>
          <Header />
          <main className={styles.content}>{children}</main>
          <div className={styles.backgroundGradient} />
        </Providers>
      </body>
    </html>
  );
}
