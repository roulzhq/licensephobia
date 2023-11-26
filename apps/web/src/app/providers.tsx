"use client";

import { queryClient } from "$api/queries";
import { QueryClientProvider } from "@tanstack/react-query";

export default function Providers({ children }) {
  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
