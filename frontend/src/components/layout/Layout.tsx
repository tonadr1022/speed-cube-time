import LeftSideBar from "./LeftSideBar";
import TopNavBar from "./TopNavBar";
import { Suspense, useEffect, useState } from "react";
import Loading from "../Loading";
import clsx from "clsx";
import { useSettings } from "../../hooks/useContext";
import { useLocation } from "react-router-dom";
import { useTheme } from "../../hooks/useTheme";

const excludeNavBarPages = ["/login", "/register"];
export default function Layout({ children }: { children: React.ReactNode }) {
  const { focusMode } = useSettings();
  const [currentURL, setCurrentURL] = useState<string | null>(null);

  const loc = useLocation();
  const showNavBar = !excludeNavBarPages.find((e) => e === loc.pathname);
  const showNav = !focusMode && showNavBar;

  useEffect(() => {
    setCurrentURL(loc?.pathname);
  }, [loc.pathname]);

  useTheme();

  return (
    <>
      <div className={clsx("flex", { "h-screen": currentURL === "/" })}>
        {showNav && (
          <div className="hidden sm:flex h-screen sticky top-0 left-0">
            <LeftSideBar />
          </div>
        )}
        <div className="w-full flex flex-col h-full">
          {showNav && (
            <div className="sm:hidden z-50 sticky top-0 left-0">
              <TopNavBar />
            </div>
          )}
          <div className="h-full">{children}</div>
        </div>
      </div>
    </>
  );
}
