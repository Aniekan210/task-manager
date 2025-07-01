/**
 * Client-side Booking Page Layout
 * Static metadata only - no dynamic placeholders
 */
export const metadata = {
    title: "Book an appointment ",
    description: "Schedule your appointment - Powered by Schedulee",

    openGraph: {
        title: "Book an appointment ",
        description: "Schedule your appointment - Powered by Schedulee",
        images: [{
            url: "https://schedulee.app/logo.png",
            width: 1200,
            height: 1200,
            alt: "Schedulee Booking",
        }],
        type: "website",
        siteName: "Schedulee",
    },

    twitter: {
        card: "summary_large_image",
        title: "Book an appointment ",
        description: "Schedule your appointment - Powered by Schedulee",
        images: "https://schedulee.app/logo.png",
    },

    robots: {
        index: false,
        follow: false,
        nocache: true,
        googleBot: {
            index: false,
            follow: false,
            noimageindex: true,
        },
    },
};

export default function BookingLayout({ children }) {
    return (
        <>
            {children}
        </>
    );
}