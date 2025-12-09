import { Container, Card, Text, Hero, UsaceBox, H2 } from "@usace/groundwork";

export default function Home() {
  const base = import.meta.env.VITE_BASE_URL;

  return (
    <>
      <Hero
        image={[
          `${base}nww-lucky-peak-dam.jpg`,
          `${base}taylorsville.jpg`,
        ]}
        alt={["Lucky Peak Dam", "Talorsville Dam"]}
        title='HH&C Shared Apps'
        subtitle='Collection of HH&C Applications'
      />

      <Container>
        <div className='mt-6'>
          <UsaceBox title='Welcome'>
            <Text>
              Welcome to the Hydrology & Hydraulics App Suite. A centralized
              workspace for the H&H CoP to manage workforce capabilities,
              cooperative stream gage funding, and project budgeting for dams,
              levees, and reservoirs. These tools help H&H practitioners and
              leaders collaborate across Districts, align with USACE Water
              Management practices, and make better-informed technical and
              financial decisions.
            </Text>
          </UsaceBox>

          <div className='grid grid-cols-1 gap-10 mt-6 sm:grid-cols-3'>
            <a href={`${base}budget`} target="_blank" rel="noopener noreferrer">
              <Card>
                <H2>Water Management Budget</H2>
                <Text>
                  Water Management Budget helps Districts build budgets for
                  their projects, tasks, and labor. Define projects, assign
                  GS-grade resources and rates, and allocate costs by business
                  lines distributed between Operations and Maintenance to align
                  with USACE Civil Works budgeting.
                </Text>
                <img
                  src={[`${base}wmbudget.png`]}
                  alt='Description'
                  className='mt-4 h-40 w-full object-cover rounded-md'
                />
              </Card>
            </a>

            <a href={`${base}gage`} target="_blank" rel="noopener noreferrer">
              <Card>
                <H2>Gage Accounting and Gage Equipment (GAGE)</H2>
                <Text>
                  The GAGE program assists Corps offices in managing funds
                  associated with the U. S. Geological Survey (USGS) Cooperative
                  Stream Gaging Program and the National Weather Service (NWS)
                  Cooperative Program.
                </Text>
                <img
                  src={[`${base}gage.png`]}
                  alt='Description'
                  className='mt-4 h-40 w-full object-cover rounded-md'
                />
              </Card>
            </a>

            <a href={`${base}workforce`} target="_blank" rel="noopener noreferrer">
              <Card>
                <H2>Water Management Workforce</H2>
                <Text>
                  Workforce is to enable senior leadership to asses the
                  strengths and weaknesses of the CoP workforce across districts
                  and divisions. Rather than utilizing various data calls with
                  spreadsheets, this will provide a single source for
                  supervisors and senior leadership to assess workforce
                  capabilities/strengths.
                </Text>
                <img
                  src={[`${base}workforce.png`]}
                  alt='Description'
                  className='mt-4 h-40 w-full object-cover rounded-md'
                />
              </Card>
            </a>
          </div>
        </div>
      </Container>
    </>
  );
}
