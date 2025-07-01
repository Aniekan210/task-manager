// app/page.js
import Link from "next/link"
import Image from "next/image"
import { Button } from "@/components/ui/button"
import { DemoButton, BackToTopButton } from "@/components/ui/interactiveButtons"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"

export default function Home() {
  const testimonials = [
    {
      quote: "schedulee.app has saved me hours each week. My clients love how easy it is to book, and I love having all my appointments in one place.",
      author: "Sarah K., Freelance Designer"
    },
    {
      quote: "The no-login feature is a game changer. My older clients were struggling with other systems, but now they book with just a phone call's worth of information.",
      author: "Michael T., Consultant"
    },
    {
      quote: "For $8.99/month, this is a no-brainer. I was paying triple for a more complex system I didn't need.",
      author: "Jessica L., Massage Therapist"
    },
    {
      quote: "Setup took 10 minutes and I was accepting bookings the same day. The trial convinced me to stay.",
      author: "David R., Tutor"
    },
    {
      quote: "My no-show rate dropped by 60% since using schedulee.app. The automated reminders are perfect.",
      author: "Emma S., Hair Stylist"
    },
    {
      quote: "Finally a booking system that doesn't overwhelm my clients with accounts and logins. Simple and effective.",
      author: "James P., Photographer"
    },
    {
      quote: "The customization options let me match my brand perfectly. Clients think it's part of my website!",
      author: "Alex M., Web Developer"
    },
    {
      quote: "I love that I can block off personal time and clients can only see my real availability. No more double bookings!",
      author: "Rachel W., Life Coach"
    }
  ]

  return (
    <div className="min-h-screen bg-white">
      {/* Navigation */}
      <nav className="container mx-auto px-6 py-4 flex justify-between items-center">
        <Link href="#" aria-label="schedulee.app Home">
          <div className="relative w-40 h-10">
            <Image
... (truncated for brevity)