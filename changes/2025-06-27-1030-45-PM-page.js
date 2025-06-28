export default async function Page({ params }) {
  const { uniqueName } = await params;
  const settings = await getBookingPageSettings(uniqueName);

  const businessName = settings["businessName"];
  const bgColor = settings["bgColor"];
  const logoUrl = settings["logoUrl"];

  const hexColor = bgColor.replace('#', '');
    
  // Parse the hex color to RGB components
  const r = parseInt(hexColor.substring(0, 2), 16) / 255;
  const g = parseInt(hexColor.substring(2, 4), 16) / 255;
  const b = parseInt(hexColor.substring(4, 6), 16) / 255;
  
  // Calculate luminance (perceived brightness)
  const luminance = 0.2126 * r + 0.7152 * g + 0.0722 * b;
  
  // Use white text if luminance is below 0.5 (threshold can be adjusted)
  const isLight = luminance < 0.5;

  

  return (
    <div>
      hi
    </div>
  );
}