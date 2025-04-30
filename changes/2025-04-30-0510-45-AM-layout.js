import "./globals.css";

export const metadata = {
  title: "Aniekanabasi Ekarika",
  description: "Aniekan Ekarika's portfolio website.",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}