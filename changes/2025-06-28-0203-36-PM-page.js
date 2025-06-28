'use client';

import Image from "next/image";
import { useEffect, useState } from "react";


export default function Page({ params }) {
  const { username } = params;
  const [settings, setSettings] = useState({
    bgColor: "#ffffff",
    logoUrl: "",
    businessName: "Aniekan's"
  });

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {

    const fetchSettings = async () => {
      try {
        const response = await fetch(`/api/getBookSettings?username=${username}`);
        const data = await response.json();
        setSettings(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchSettings();
  }, []);


  const businessName = settings['businessName'];
  const bgColor = settings['bgColor'];
  const logoUrl = settings['logoUrl'];

  const hexColor = bgColor.replace('#', '');

  // Parse hex to RGB (0-255)
  const r = parseInt(hexColor.substring(0, 2), 16);
  const g = parseInt(hexColor.substring(2, 4), 16);
  const b = parseInt(hexColor.substring(4, 6), 16);

  // Convert RGB to HSL
  const r1 = r / 255, g1 = g / 255, b1 = b / 255;
  const max = Math.max(r1, g1, b1), min = Math.min(r1, g1, b1);
  let h, s, l = (max + min) / 2;
... (truncated for brevity)