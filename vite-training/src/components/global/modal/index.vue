<script setup lang="ts">
import {
  Dialog,
  DialogPanel,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import { Icon } from "@iconify/vue";

type ModalProps = {
  customClass?: string;
  state: boolean;
  customClassOverlay?: string;
  zIndex?: string;
};

withDefaults(defineProps<ModalProps>(), {
  zIndex: "z-[1000]"
});

const emit = defineEmits(["close"]);

const closeModal = () => {
  emit("close");
};
</script>

<template>
  <TransitionRoot appear :show="state" as="template">
    <Dialog as="div" @click="closeModal" class="relative" :class="zIndex">
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

      <div
        class="fixed font-inter inset-0 overflow-y-auto"
        :style="customClassOverlay"
      >
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
            <DialogPanel
              class="p-4 overflow-hidden text-left align-middle transition-all transform bg-white rounded-3xl shadow-xl"
              :class="customClass"
            >
              <div class="relative flex items-center justify-end">
                <Icon
                  icon="iconamoon:close"
                  class="w-6 h-6 text-gray-500 cursor-pointer"
                  @click="closeModal"
                />
              </div>
              <slot name="modal"></slot>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
