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
      cell: ({row}) => h('div', {class: ''}, row.getValue(key)),
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


</script>

<template>
  <div class="container" v-if="isLoaded">

    <div class="flex gap-2 items-center py-4">
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

    <div class="rounded-md border w-min mx-auto">
      <Table>
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
  </div>
</template>

<style scoped>

</style>