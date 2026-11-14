import { describe, expect, test } from "vitest";
import { render } from "@testing-library/react";
import { Header } from "./header";

describe("header", () => {
  test("should match snapshot", () => {
    const { asFragment } = render(<Header />);
    expect(asFragment()).toMatchSnapshot();
  });
});
