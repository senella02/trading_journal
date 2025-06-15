'use client'
import { useEffect } from "react";
import AddTradeForm from "../components/ui/form/AddTradeForm";
import { useUser} from "@/context/userContext";
export default function AddTradePage() {
    const {hardWireUser} = useUser();
    useEffect(() => {
        hardWireUser();
    }, []);
    return (
        <div>
            Landing page after login
        </div>
    );
}