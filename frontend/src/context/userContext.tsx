"use client";
import { createContext, useContext, useState, ReactNode, useEffect } from "react";
import Cookies from "js-cookie";

export interface User {
    userId: string;
    userName: string;
    token: string;
}

interface UserContextType {
    user: User | null;
    setUser: (user: User | null) => void;
    logout: () => void;
    hardWireUser: () => void; //Define what our context will provide when we use it! So client can use it from useUser()
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider = ({ children} : {children: ReactNode}) => {
    const [user, setUser] = useState<User | null>(null);
    //always mounting on loading page to get user data, but only do once when we enter so we use the useEffect
    useEffect(() => {
        const token = Cookies.get("token");
        const userId = Cookies.get("userId");
        const userName = Cookies.get("userName");

        if(token && userId && userName) {
            setUser({userId, userName, token})
        }
    }, []);

    
    const handleLoginSuccess = (user: { token: string; userId: string; userName: string }) => {
        Cookies.set("token", user.token);
        Cookies.set("userId", user.userId);
        Cookies.set("userName", user.userName);

        setUser(user); // set context
    };
    const hardWireUser = () => {
        setUser({ userId: "id0101", userName: "thananan", token: "token1234" });
    }
    const logout = () => {
        setUser(null);
        Cookies.remove("token");
        Cookies.remove("userId");
        Cookies.remove("username");
        // optionally redirect
    };

    return (
    <UserContext.Provider value={{ user, setUser, logout, hardWireUser }}>
      {children}
    </UserContext.Provider>
  );
}

export const useUser = () => {
  const ctx = useContext(UserContext); // <-- this line
  if (!ctx) throw new Error("useUser must be used inside a UserProvider");
  return ctx;
};
