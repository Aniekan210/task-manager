import "./styles.css"

export default async function Page({ params }) {
  const { uniqueName } = await params;
  const settings = await getBookingPageSettings(uniqueName);
  return (
    <div>
      hi
    </div>
  );
}