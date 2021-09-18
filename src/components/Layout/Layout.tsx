import { Container } from "@chakra-ui/layout";
import React from "react";
import Navbar from "./Navbar";

export default function Layout({ children }: any) {
  return (
    <Container maxWidth="8xl">
      <Navbar />
      {children}
    </Container>
  );
}
