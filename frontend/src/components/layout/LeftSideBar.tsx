import clsx from "clsx";
import React from "react";
import {
  FaChevronLeft,
  FaChevronRight,
  FaGear,
  FaTable,
} from "react-icons/fa6";
import { RiLogoutBoxLine } from "react-icons/ri";
import { RxTimer } from "react-icons/rx";
import {
  useAuth,
  useLayoutContext,
  useOnlineContext,
} from "../../hooks/useContext";
import { Link, useLocation, useNavigate } from "react-router-dom";

type Props = {
  children: React.ReactNode;
};

const LeftSideBarShell = ({ children }: Props) => {
  const { navCollapsed, setNavCollapsed } = useLayoutContext();
  const auth = useAuth();
  const { online } = useOnlineContext();
  const navigate = useNavigate();
  return (
    <nav
      className={clsx(
        "h-screen flex flex-col bg-base-300 w-16",
        !navCollapsed ? "md:w-56" : "md:w-16",
      )}
    >
      <div className="flex pt-4 justify-center">
        <img
          src={"/pwa-512x512.png"}
          alt="ChronoCube Logo"
          width={30}
          height={30}
        />
        {!navCollapsed && (
          <span className="hidden md:flex pl-2 text-xl font-bold self-center">
            CubeChron
          </span>
        )}
      </div>
      <ul
        className={clsx(
          "flex-1 pt-2 text-base-content flex flex-col items-center md:items-start",
          navCollapsed && "md:items-center md:ml-0",
          !navCollapsed && "md:ml-4",
        )}
      >
        {children}
      </ul>
      <div className="flex justify-center mb-10 gap-y-8">
        {auth.user && online && (
          <button
            className="absolute bottom-24 bg-base-300 outline-none btn btn-sm btn-neutral-focus p-1 m-0 rounded-full"
            onClick={() => auth.logout(() => navigate("/login"))}
          >
            <RiLogoutBoxLine className="w-6 h-6" />
          </button>
        )}
        <button
          className=" opacity-75 hover:opacity-100 transition-opacity cursor-pointer absolute bottom-12 h-8"
          onClick={() => setNavCollapsed(!navCollapsed)}
        >
          {navCollapsed ? (
            <FaChevronRight />
          ) : (
            <div className="hidden md:flex items-center">
              <FaChevronLeft />
              <span className="pl-2">Collapse</span>
            </div>
          )}
        </button>
      </div>
    </nav>
  );
};

const LeftSideBar = () => {
  const { navCollapsed } = useLayoutContext();
  const { pathname } = useLocation();
  return (
    <LeftSideBarShell>
      <SideBarItem
        navCollapsed={navCollapsed}
        icon={<RxTimer />}
        text="Timer"
        href="/"
        active={pathname === "/"}
      />
      {/* <SideBarItem */}
      {/*   navCollapsed={navCollapsed} */}
      {/*   icon={<FaChartPie />} */}
      {/*   text="Stats" */}
      {/*   href="/stats" */}
      {/* /> */}
      {/* <SideBarItem */}
      {/*   navCollapsed={navCollapsed} */}
      {/*   icon={<FaCubesStacked />} */}
      {/*   text="Solves" */}
      {/*   href="/solves" */}
      {/* /> */}
      <SideBarItem
        navCollapsed={navCollapsed}
        icon={<FaTable />}
        text="Sessions"
        href="/sessions"
        active={pathname === "/sessions"}
      />
      <SideBarItem
        navCollapsed={navCollapsed}
        icon={<FaGear />}
        text="Settings"
        href="/settings"
        active={pathname === "/settings"}
      />
    </LeftSideBarShell>
  );
};

export default LeftSideBar;

type SideBarItemProps = {
  icon: React.ReactNode;
  text: string;
  href: string;
  navCollapsed: boolean;
  active: boolean;
};

export const SideBarItem = ({
  navCollapsed,
  icon,
  text,
  href,
  active,
}: SideBarItemProps) => {
  return (
    <li className="py-2">
      <Link
        to={href}
        className={clsx(
          "flex items-center hover:opacity-100 transition-opacity cursor-pointer",
          active ? "opacity-100" : "opacity-50",
        )}
      >
        {icon}
        {!navCollapsed && (
          <>
            <div className="hidden md:block">
              <span className="pl-2">{text}</span>
            </div>
          </>
        )}
      </Link>
    </li>
  );
};
