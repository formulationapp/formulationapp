<script setup lang="ts">

import TableBody from "@/components/ui/table/TableBody.vue";
import {
  ColumnFiltersState,
  FlexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useVueTable,
  VisibilityState
} from "@tanstack/vue-table";
import TableHead from "@/components/ui/table/TableHead.vue";
import TableRow from "@/components/ui/table/TableRow.vue";
import TableHeader from "@/components/ui/table/TableHeader.vue";
import TableCell from "@/components/ui/table/TableCell.vue";
import {h, onMounted, ref} from "vue";
import {valueUpdater} from "@/lib/utils";
import {ArrowUpDown} from "lucide-vue-next";
import {Button} from "@/components/ui/button";
import DropdownMenuTrigger from "@/components/ui/dropdown-menu/DropdownMenuTrigger.vue";
import DropdownMenu from "@/components/ui/dropdown-menu/DropdownMenu.vue";
import DropdownMenuContent from "@/components/ui/dropdown-menu/DropdownMenuContent.vue";
import DropdownMenuCheckboxItem from "@/components/ui/dropdown-menu/DropdownMenuCheckboxItem.vue";
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";
import AlertTitle from "@/components/ui/alert/AlertTitle.vue";
import Alert from "@/components/ui/alert/Alert.vue";
import AlertDescription from "@/components/ui/alert/AlertDescription.vue";
import Input from "@/components/ui/input/Input.vue";
import {copyToClipboard} from "@/utils";

const sorting = ref<SortingState>([])
const columnFilters = ref<ColumnFiltersState>([])
const columnVisibility = ref<VisibilityState>({})
const rowSelection = ref({})

const forms = useForms();
const route = useRoute();

let data = [];
let columns = [];
let aliases = {};
let table = null;
const isLoaded = ref(false);


onMounted(async () => {
  await forms.load(parseInt(route.params.workspaceID));
  forms.select(parseInt(route.params.formID));
  data = await forms.loadSubmissions();

  aliases = forms.getAliases();

  columns = Object.keys(aliases).map(key => {
    return {
      accessorKey: key,
      header: ({column}) => {
        return h(Button, {
          variant: 'ghost',
          onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
        }, () => [aliases[key], h(ArrowUpDown, {class: 'ml-2 h-4 w-4'})])
      },
      cell: ({row}) => {
        const value = row.getValue(key);
        if (value != undefined && value.startsWith('{') && value.endsWith('}')) {
          const selected = JSON.parse(value);
          const badges = Object.keys(selected).filter(choice => selected[choice]).map(choice => h('span', {class: 'bg-blue-100 text-blue-800 text-xs font-medium mr-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300 w-min'}, choice));
          return h('div', {class: ''}, badges);
        }
        return h('div', {class: ''}, row.getValue(key));
      },
    };
  });

  table = useVueTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onSortingChange: updaterOrValue => valueUpdater(updaterOrValue, sorting),
    onColumnFiltersChange: updaterOrValue => valueUpdater(updaterOrValue, columnFilters),
    onColumnVisibilityChange: updaterOrValue => valueUpdater(updaterOrValue, columnVisibility),
    onRowSelectionChange: updaterOrValue => valueUpdater(updaterOrValue, rowSelection),
    state: {
      get sorting() {
        return sorting.value
      },
      get columnFilters() {
        return columnFilters.value
      },
      get columnVisibility() {
        return columnVisibility.value
      },
      get rowSelection() {
        return rowSelection.value
      },
    },
  });
  isLoaded.value = true;
});

const link = ref('http://' + window.location.host + '/f/' + forms.form.secret);

</script>

<template>
  <div class="container" v-if="isLoaded">

    <div v-if="table.getRowModel().rows?.length">
      <div class="flex gap-2 items-center py-4 w-3/4" >
        <!--      <Input-->
        <!--          class="max-w-sm"-->
        <!--          placeholder="Filter submissions..."-->
        <!--          :model-value="table.getColumn('email')?.getFilterValue() as string"-->
        <!--          @update:model-value=" table.getColumn('email')?.setFilterValue($event)"-->
        <!--      />-->
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button variant="outline" class="ml-auto">
              Columns
              <v-icon name="bi-chevron-down" class="ml-2 h-4 w-4"/>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuCheckboxItem
                v-for="column in table.getAllColumns().filter((column) => column.getCanHide())"
                :key="column.id"
                class="capitalize"
                :checked="column.getIsVisible()"
                @update:checked="(value) => {
              column.toggleVisibility(value)
            }"
            >
              {{ aliases[column.id.toLowerCase()] }}
            </DropdownMenuCheckboxItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <div class="rounded-md border w-1/2 mx-auto " >
        <Table class="w-full">
          <TableHeader>
            <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
              <TableHead v-for="header in headerGroup.headers" :key="header.id">
                <FlexRender v-if="!header.isPlaceholder" :render="header.column.columnDef.header"
                            :props="header.getContext()"/>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <template v-if="table.getRowModel().rows?.length">
              <TableRow
                  v-for="row in table.getRowModel().rows"
                  :key="row.id"
                  :data-state="row.getIsSelected() && 'selected'"
              >
                <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                  <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()"/>
                </TableCell>
              </TableRow>
            </template>

            <TableRow v-else>
              <TableCell
                  col-span="{columns.length}"
                  class="h-24 text-center"
              >
                No results.
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <div class="flex items-center justify-end space-x-2 py-4  w-1/2 mx-auto">
        <div class="flex-1 text-sm text-muted-foreground">
          {{ table.getFilteredSelectedRowModel().rows.length }} of
          {{ table.getFilteredRowModel().rows.length }} row(s) selected.
        </div>
        <div class="space-x-2">
          <Button
              variant="outline"
              size="sm"
              :disabled="!table.getCanPreviousPage()"
              @click="table.previousPage()"
          >
            Previous
          </Button>
          <Button
              variant="outline"
              size="sm"
              :disabled="!table.getCanNextPage()"
              @click="table.nextPage()"
          >
            Next
          </Button>
        </div>
      </div>
    </div>

    <div v-else class="w-1/2 mx-auto mt-12">
      <Alert class="mb-4">
        <AlertTitle>Oh no!</AlertTitle>
        <AlertDescription>
          There are no submissions for this form yet. Share this form with your users to start collecting submissions.
        </AlertDescription>
      </Alert>
      <Input v-model="link" readonly @click="copyToClipboard(link)" class="cursor-pointer"></Input>
    </div>
  </div>
</template>

<style scoped>

</style>