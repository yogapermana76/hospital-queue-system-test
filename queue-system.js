const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

let patientQueue = [];
let roundRobinMode = false;
let roundRobinGender = 'F'; // Initialize with Female

function isMRNumberUnique(mrNumber) {
  return !patientQueue.some((patient) => patient.mrNumber === mrNumber);
}

function printQueue() {
  console.log('Queue:');
  for (const patient of patientQueue) {
    console.log(`${patient.mrNumber} ${patient.gender}`);
  }
}

function processCommand(command) {
  const args = command.split(' ');

  if (args[0] === 'IN') {
    const mrNumber = args[1];
    const gender = args[2];

    if (!mrNumber || !gender) {
      console.log(
        'error: Missing arguments, please input for example IN MR1234 M'
      );
      return;
    }

    if (gender !== 'M' && gender !== 'F') {
      console.log('error: please input gender as M or F');
      return;
    }

    if (!mrNumber.match(/^MR\d{4}$/)) {
      console.log('error: Invalid MRNumber format');
      return;
    }

    if (!isMRNumberUnique(mrNumber)) {
      console.log(`error: Patient with ${mrNumber} already in queue`);
      return;
    }

    patientQueue.push({ mrNumber: mrNumber, gender: gender });
    console.log(`Added ${mrNumber} ${gender} to the queue.`);
  } else if (args[0] === 'OUT') {
    if (patientQueue.length === 0) {
      console.log('Queue is empty.');
      return;
    }

    const patient = roundRobinMode
      ? patientQueue.find((p) => p.gender === roundRobinGender)
      : patientQueue.shift();

    patientQueue = patientQueue.filter((p) => p.mrNumber !== patient.mrNumber);

    if (!patient) {
      roundRobinMode = false; // Exit round-robin mode if no eligible patient is found.
      console.log(
        'No eligible patient found for round-robin. Switching to default.'
      );
      return;
    }

    console.log(`send: ${patient.mrNumber} ${patient.gender} > OUT`);
  } else if (args[0] === 'ROUNDROBIN') {
    roundRobinMode = true;
    console.log('Switched to round-robin mode.');
  } else if (args[0] === 'DEFAULT') {
    roundRobinMode = false;
    console.log('Switched to default mode.');
  } else if (args[0] === 'EXIT') {
    rl.close();
  } else {
    console.log('error: Invalid command');
  }
}

function main() {
  console.log('Welcome to the Hospital Patient Queue System!');
  console.log('Commands: IN, OUT, ROUNDROBIN, DEFAULT, EXIT');

  rl.on('line', (input) => {
    processCommand(input);
    printQueue();
  });
}

main();
