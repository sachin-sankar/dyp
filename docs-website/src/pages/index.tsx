import type {ReactNode} from 'react';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className="hero hero--primary">
      <div className="container">
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div style={{display: 'flex', gap: '1rem', justifyContent: 'center', marginTop: '2rem'}}>
          <Link className="button button--secondary button--lg" to="/docs/intro">
            Get Started
          </Link>
          <Link className="button button--outline button--lg" to="https://github.com/sachin-sankar/dyp">
            GitHub
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title="Home"
      description={siteConfig.tagline}>
      <HomepageHeader />
      <main>
        <div className="container" style={{padding: '3rem 0', textAlign: 'center'}}>
          <h2>Quick Links</h2>
          <div style={{display: 'flex', gap: '2rem', justifyContent: 'center', marginTop: '1rem', flexWrap: 'wrap'}}>
            <Link to="/docs/installation" className="card">
              <h3>Installation</h3>
              <p>Set up DYP on your system</p>
            </Link>
            <Link to="/docs/prompt-syntax" className="card">
              <h3>Prompt Syntax</h3>
              <p>Learn how to write prompts</p>
            </Link>
            <Link to="/docs/examples" className="card">
              <h3>Examples</h3>
              <p>Browse example templates</p>
            </Link>
          </div>
        </div>
      </main>
    </Layout>
  );
}