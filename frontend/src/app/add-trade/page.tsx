import AddTradeForm from "../../components/ui/form/AddTradeForm";

export default function AddTradePage() {
    return (
         <div className="flex h-screen overflow-hidden">
            {/* Your sidebar should already be rendered outside this or as fixed */}

            <main className="flex flex-grow flex-1 overflow-hidden">
                {/* Form area */}
                <div className="flex-grow min-w-0 overflow-y-auto p-4 ">
                    <AddTradeForm />
                </div>

                {/* Chart area */}
                <div className="w-[900px] flex-shrink-0 bg-amber-200 p-4 overflow-y-auto">
                <p>Chart placeholder</p>
                </div>
            </main>
        </div>
    );
}