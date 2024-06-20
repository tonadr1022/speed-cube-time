import React from "react";

type Props = {
  title: string;
  children: React.ReactNode;
};

const PageWrapper = ({ title, children }: Props) => {
  return (
    <div className="p-6 max-w-md">
      <h1 className="text-6xl font-semibold">{title}</h1>
      {children}
    </div>
  );
};

export default PageWrapper;
