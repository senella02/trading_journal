"use client"

import { useSidebar } from "@/components/ui/sidebar/sidebar"
import { AppSidebar } from "@/components/app-sidebar"
import { SidebarInset } from "@/components/ui/sidebar/sidebar"
import { SidebarTrigger } from "@/components/ui/sidebar/sidebar"
import { Separator } from "@/components/ui/separator"
import { DynamicBreadcrumb } from "@/components/dynamic-breadcrumb"

export default function LayoutWrapper({ children }: { children: React.ReactNode }) {
  const {
    state
  } = useSidebar()

  return (
     <div className="flex h-screen w-screen overflow-hidden">
        {/* SIDEBAR LEFT */}
        <div className={`${
          state === "collapsed" ? "w-[--sidebar-width-mobile]" : "w-[--sidebar-width]"
        } shrink-0 transition-all duration-300 border-r bg-sidebar z-50`}>
          <AppSidebar />
        </div>

        {/* MAIN CONTENT */}
        <div className="flex flex-col flex-1 overflow-hidden">
          <SidebarInset>
            <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
              <SidebarTrigger className="-ml-1 text-content-main" />
              <Separator
                orientation="vertical"
                className="mr-2 data-[orientation=vertical]:h-4"
              />
              <DynamicBreadcrumb />
            </header>
            <div className="bg-background flex-1 overflow-auto max-w-full">
              {children}
            </div>
          </SidebarInset>
        </div>
      </div>
  )
}
