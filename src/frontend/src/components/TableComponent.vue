<script setup lang="ts"></script>
<template>
  <div class="w-full">
    <table>
      <tr class="text-darkgray font-bold">
        <th class="hidden sm:table-cell">Id</th>
        <th>Name</th>
        <th class="hidden sm:table-cell">Email</th>
        <th>Guests</th>
        <th>Time Slot</th>
      </tr>
      <tr>
        <td class="hidden sm:table-cell">20919</td>
        <td>Maria Anders</td>
        <td class="hidden sm:table-cell">maria@gmail.com</td>
        <td>5</td>
        <td>12:00 - 13:30</td>
      </tr>
    </table>
  </div>
</template>
<style>
table {
  width: 100%;
}

td,
th {
  text-align: center;
  padding: 8px;
  max-width: 20%;
}
</style>
