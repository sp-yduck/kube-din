import * as React from 'react';
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import axios from "axios";
import Title from './Title';

function preventDefault(event: React.MouseEvent) {
  event.preventDefault();
}

export default function Jobs() {
  const [jobs, setJobs] = React.useState([]);

  React.useEffect(() => {
    const setResponse = async () => {
      // to do : use port number from env var
      const res = await axios.get("http://localhost:1234/api/job");
      console.log(res)
      setJobs(res.data.items)
    };
    setResponse();
  }, []);

  return (
    <React.Fragment>
      <Title>Recent Jobs</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Creation Time</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>NameSpace</TableCell>
            <TableCell>Succeeded</TableCell>
            <TableCell align="right">Completion</TableCell>
          </TableRow>
        </TableHead>
          <TableBody>
            {jobs.map((job: any) => (
              <TableRow key={job.metadata.uid}>
                <TableCell>{job.metadata.creationTimestamp}</TableCell>
                <TableCell>{job.metadata.name}</TableCell>
                <TableCell>{job.metadata.namespace}</TableCell>
                <TableCell>{job.status.succeeded}</TableCell>
                <TableCell align="right">{job.spec.completions}</TableCell>
              </TableRow>
            ))}
          </TableBody>
      </Table>
      {/* WIP */}
      {/* <Link color="primary" href="#" onClick={preventDefault} sx={{ mt: 3 }}>
        See more jobs
      </Link> */}
    </React.Fragment>
  );
}
