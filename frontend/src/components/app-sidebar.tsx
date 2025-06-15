"use client";
import * as React from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { SearchForm } from "@/components/search-form";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar/sidebar";
import { Home } from "lucide-react";
// Define sidebar data
const data = {
  versions: ["1.0.1", "1.1.0-alpha", "2.0.0-beta1"],
  navMain: [
    {
      title: "Trade management",
      url: "#",
      items: [
        {
          title: "Add Trade",
          url: "/add-trade",
        },
        {
          title: "Playbook",
          url: "/playbook",
        },
      ],
    },
    {
      title: "Analysis",
      url: "#",
      items: [
        {
          title: "Dashboard",
          url: "/dashboard",
        },
        {
          title: "Trade Tracking",
          url: "/trade-tracking",
        },
      ],
    },
    {
      title: "Others",
      url: "#",
      items: [
        {
          title: "Setting",
          url: "/settings",
        },
        {
          title: "FAQ",
          url: "/faq",
        },
        {
          title: "About Us",
          url: "/about-us",
        },
      ],
    },
  ],
};

//Define type for item
type SidebarItem = {
  parent: string;
  title: string;
  url: string;
};

const SidebarSelectionContext = React.createContext<{
  selectedItem: SidebarItem | null;
  setSelectedItem: (item: SidebarItem) => void;
}>({
  selectedItem: null,
  setSelectedItem: () => {}, //place holder function
});

export const SideBarSelectionProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [selectedItem, setSelectedItem] = React.useState<SidebarItem | null>(
    null
  );

  return (
    <SidebarSelectionContext.Provider value={{ selectedItem, setSelectedItem }}>
      {children}
    </SidebarSelectionContext.Provider>
  );
};

export const useSidebarSelection = () => {
  const context = React.useContext(SidebarSelectionContext);
  if (!context) {
    throw new Error(
      "useSidebarSelection must be used within SidebarSelectionProvider"
    );
  }
  return context;
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const router = useRouter();
  const { setSelectedItem, selectedItem } = useSidebarSelection();
  return (
    <Sidebar {...props}>
      <SidebarHeader>

          <Link href="/" className="inline-flex items-center gap-2 p-4 text-content-main hover:text-content-grey">
            <Home className="w-7 h-7"></Home>
            <span >
              Trading Journal
          </span>
          </Link>

        <SearchForm />
      </SidebarHeader>
      <SidebarContent>
        {/* We create a SidebarGroup for each parent. */}
        {data.navMain.map((parent) => (
          <SidebarGroup key={parent.title}>
            <SidebarGroupLabel>{parent.title}</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {parent.items.map((child) => (
                  <SidebarMenuItem key={child.title}>
                    <SidebarMenuButton asChild isActive={selectedItem?.url === child.url}>
                      <a
                        href={child.url}
                        onClick={(e) => {
                          console.log(child.url);
                          e.preventDefault(); // prevent full page reload
                          setSelectedItem({
                            parent: parent.title,
                            title: child.title,
                            url: child.url,
                          });
                          router.push(child.url); //navigate to page
                        }}
                      >
                        {child.title}
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  );
}
