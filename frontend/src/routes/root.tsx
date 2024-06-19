import { Outlet } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Layout from "../components/layout/Layout";

export default function Root() {
  return (
    <Layout>
      <Outlet />
      <ToastContainer />
    </Layout>
  );
}
