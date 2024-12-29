describe("login with correct credentials", () => {
  it("logging in with the right credentials should redirect user to the home page", () => {
    cy.visit("http://localhost:5173");
    cy.get("#username").type("kavish");
    cy.get("#login").click();
    cy.contains("Search");
    cy.setCookie("auth_token", "");
  });
});

describe("login with incorrect credentials", () => {
  it("logging in with the incorrect credentials should alert the user", () => {
    cy.visit("http://localhost:5173");
    cy.get("#username").type("kavish");
    cy.get("#login").click();
    cy.on("window:alert", (string) =>
      expect(string).to.equal("Incorrect Credentials.")
    );
  });
});
