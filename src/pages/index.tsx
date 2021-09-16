import type { NextPage } from "next";
import { Container, Heading } from "@chakra-ui/layout";

const Home: NextPage = () => {
  return (
    <Container maxWidth="container.xl">
      <Heading
        textAlign="center"
        fontWeight="light"
        fontSize="7xl"
        marginY="24"
      >
        The <strong>best</strong> place to display your{" "}
        <a style={{ color: "#AB41FF" }}>
          <strong>art</strong>
        </a>{" "}
        and manage comissions!
      </Heading>
    </Container>
  );
};

export default Home;
