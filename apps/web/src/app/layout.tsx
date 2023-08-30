import "@fontsource/ibm-plex-mono";
import "$styles/base.scss";

import Header from "$components/header";
import styles from "./layout.module.scss";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}): JSX.Element {
  return (
    <html lang="en">
      <body>
        <Header />
        <main className={styles.content}>{children}</main>
        <div className={styles.backgroundGradient} />
      </body>
    </html>
  );
}
