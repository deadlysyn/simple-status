import React from "react";

const NavBar = () => {
  return (
    <nav
      className="navbar is-transparent"
      role="navigation"
      aria-label="main navigation"
    >
      <div className="navbar-brand">
        <a className="navbar-item" href="https://sonatype.com">
          <img
            src="https://www.sonatype.com/hs-fs/hubfs/SON_logo_main@2x%20copy%20trimmed.png?width=165&name=SON_logo_main@2x%20copy%20trimmed.png"
            alt="Sonatype: Automating DevSecOps"
            width="161"
            height="28"
          />
        </a>
      </div>

      <div className="navbar-menu">
        <div className="navbar-start">
          <a className="navbar-item" href="https://sonatype.com">
            Home
          </a>

          <div className="navbar-item has-dropdown is-hoverable">
            <a className="navbar-link" href="https://help.sonatype.com/docs">
              Resources
            </a>
            <div className="navbar-dropdown is-boxed">
              <a className="navbar-item" href="https://help.sonatype.com/docs">
                Overview
              </a>
              <a
                className="navbar-item"
                href="https://help.sonatype.com/repomanager2"
              >
                Repository Manager v2
              </a>
              <a
                className="navbar-item"
                href="https://help.sonatype.com/repomanager3"
              >
                Repository Manager v3
              </a>
              <a
                className="navbar-item"
                href="https://help.sonatype.com/iqserver"
              >
                IQ Server
              </a>

              <hr className="navbar-divider" />
              <a className="navbar-item" href="https://status.maven.org">
                Central Status
              </a>
            </div>
          </div>
        </div>
      </div>

      <div className="navbar-end">
        <div className="navbar-item">
          <div className="field is-grouped">
            <p className="control">
              <a
                className="bd-tw-button button is-info"
                data-social-network="Twitter"
                data-social-action="tweet"
                data-social-target="https://sonatype.com"
                target="_blank"
                rel="noopener noreferrer"
                href="https://twitter.com/intent/tweet?text=Service Status&amp;hashtags=sonatype_ops&amp;url=https://sonatype.com&amp;via=status"
              >
                <span className="icon">
                  <i className="fab fa-twitter"></i>
                </span>
                <span>Tweet</span>
              </a>
            </p>
            <p className="control">
              <a
                className="button is-info"
                href="https://github.com/jgthms/bulma/releases/download/0.9.0/bulma-0.9.0.zip"
              >
                <span className="icon">
                  <i className="fas fa-download"></i>
                </span>
                <span>Download</span>
              </a>
            </p>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
