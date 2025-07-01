"use client";

import { createContext, useContext, useState, useEffect } from "react";
import { useRouter } from "next/navigation";

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const router = useRouter();
  const [authData, setAuthData] = useState({
    user: null,
    hasPaid: false,
    daysLeft: 12,
  });

  useEffect(() => {
    const authCheck = async () => {
      try {
        const res = await fetch("/api/authenticate");
        const data = await res.json();

        if (data.error) {
          router.push("/signup");
          return;
        }

        if (!data.hasPaid && data.daysLeft === 0) {
          router.push("/upgrade");
          return;
        }

        setAuthData({
          ...data,
          isLoading: false,
        });
      } catch (error) {
        router.push("/signup");
      }
    };

    authCheck();
  }, [router]);

  return (
    <AuthContext.Provider value={authData}>{children}</AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
... (truncated for brevity)