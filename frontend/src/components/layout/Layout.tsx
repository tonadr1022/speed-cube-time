import LeftSideBar from "./LeftSideBar";
import TopNavBar from "./TopNavBar";
import { useEffect, useState } from "react";
import clsx from "clsx";
import { useSettings } from "../../hooks/useContext";
import { useLocation } from "react-router-dom";
import { useTheme } from "../../hooks/useTheme";
import FooterNav from "./FooterNav";

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
      <div
        className={clsx("flex h-screen", { "h-screen": currentURL === "/" })}
      >
        {showNav && (
          <div className="hidden md:flex h-screen sticky top-0 left-0">
            <LeftSideBar />
          </div>
        )}
        <div className="w-full flex flex-col h-full">
          {showNav && (
            <TopNavBar className="md:hidden z-50 sticky top-0 left-0" />
          )}
          <div className="flex flex-col h-full w-full">
            <div className="flex-1">{children}</div>
            <FooterNav className="md:hidden" />
          </div>
        </div>
      </div>
    </>
  );
}
