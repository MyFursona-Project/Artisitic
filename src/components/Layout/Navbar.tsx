import { Image } from "@chakra-ui/image";
import { Flex, Box, Link } from "@chakra-ui/layout";
import React from "react";
import { Input } from "@chakra-ui/react";

export default function Navbar() {
  return (
    <Flex
      marginBottom="7"
      padding="2"
      justifyContent="space-around"
      alignItems="center"
    >
      <Image src="/images/ArtisticallyDark.svg" alt="Artistically Logo" />
      <Input
        placeholder="Search for Tags, Artist, Topics..."
        background="white"
        size="sm"
        borderRadius="50"
        width="55%"
      />
      <Flex align="center">
        <Link href="/" marginX="20px" fontWeight="bold">
          <a>Home</a>
        </Link>
        <Link href="/community" marginX="20px" fontWeight="bold">
          <a>Community</a>
        </Link>
        <Link href="/about" marginX="20px" fontWeight="bold">
          <a>About Us</a>
        </Link>
        <Box
          background="#AB41FF"
          marginX="20px"
          paddingX="7"
          paddingY="2"
          borderRadius="50"
        >
          Login
        </Box>
      </Flex>
    </Flex>
  );
}
