<script setup lang="ts">
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogPanel,
  DialogTitle,
} from "@headlessui/vue";
import { ref, computed } from "vue";
// import { Icon } from "@iconify/vue";
import Button from "../button/index.vue";
// import Nodata from "@/components/general/nodata.vue";

const emit = defineEmits<(event: "click") => void>();

const props = defineProps({
  title: String,
  question: String,
  technical_note: String,
  keyword: String as any,
  // answer: String,
  fix_solution: String,
  workaround: String,
});

const isOpen = ref(false);

function closeModal() {
  isOpen.value = false;
}

function openModal() {
  isOpen.value = true;
}

</script>

<template>
  <Button
    class="text-gray-600"
    size="md"
    text="Preview Article"
    type="secondary"
    icon="iconamoon:eye"
    @click="openModal"
  />
  <TransitionRoot appear :show="isOpen" as="template">
    <Dialog as="div" @click.self="closeModal" class="relative z-[500]">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="duration-200 ease-in"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div
          class="flex items-center justify-center min-h-full p-4 text-center"
        >
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
          woi
          <div class="flex justify-end">
                  <Button
                    class="text-gray-600"
                    size="md"
                    text="Cancel"
                    type="secondary"
                    @click="closeModal"
                  />
                </div>
            <DialogPanel
              class="w-full max-w-6xl p-4 overflow-hidden text-left align-middle transition-all transform bg-white rounded-lg shadow-xl"
            >
              
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
