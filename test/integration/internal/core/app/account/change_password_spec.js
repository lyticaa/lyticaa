describe('Change Password', () => {
    it('Requests to reset a users password', () => {
        cy.visit("/")
        cy.url().should('include', 'l.authfor.me/login')
        cy.get('input[name="email"]').type(Cypress.env("USER"))
        cy.get('input[name="password"]').type(Cypress.env("PASSWORD"))
        cy.get('span.auth0-label-submit').click()
    });
});
