<script setup lang="ts">
import { computed, toRefs } from "vue";
import { useField, useIsFieldTouched } from "vee-validate";
import { Icon } from "@iconify/vue";

type TextProps = {
  value?: string;
  name: string;
  label?: string;
  size?: "md" | "lg" | "xl";
  placeholder?: string;
  icon?: string;
  disabled?: boolean;
  readonly?: boolean;
  primary?: boolean;
};

const props = defineProps<TextProps>();

const selectedSize = computed(() => {
  switch (props.size) {
    case "md":
      return "text-md px-4 py-2.5 rounded-lg font-normal";
    case "lg":
      return "px-5 py-3 rounded-2xl font-semibold";
    case "xl":
      return "px-5 py-3 rounded-lg font-semibold text-lg";
    default:
      break;
  }
});

const emit = defineEmits<(event: "update:value", value: string) => void>();

const {
  value: inputValue,
  errorMessage,
  handleChange,
  meta,
} = useField(toRefs(props).name);
const isTouched = useIsFieldTouched(toRefs(props).name);

const modelValue = computed({
  get() {
    return (inputValue.value ? inputValue.value : null) as string;
  },
  set(value: string) {
    handleChange(value);
    emit("update:value", value);
  },
});
</script>

<template>
  <div
    class="w-full font-normal last:mb-0 font-inter"
    :class="{ success: meta.valid }"
  >
    <label
      v-if="props.label"
      :for="props.name"
      class="inline-flex items-center gap-1 mb-1 form-label dark:text-white"
      :class="{
        'text-red-500 font-bold': !!errorMessage && isTouched,
        selectedSize,
      }"
    >
      {{ props.label }}
      <span v-if="props.primary">
        <span class="font-bold text-red-500">*</span>
      </span>
    </label>
    <div class="relative">
      <input
        :id="props.name"
        :name="props.name"
        type="password"
        class="w-full p-2 border rounded-lg border-slate-300 h-11 focus:border-slate-300 focus:outline-none focus:ring-0 disabled:bg-slate-50 disabled:cursor-not-allowed dark:bg-slate-700 dark:border-slate-600"
        :class="{
          'border-red-500 dark:border-red-500': !!errorMessage && isTouched,
          'pl-10': !!props.icon,
          selectedSize,
        }"
        :disabled="props.disabled"
        :readonly="props.readonly"
        v-model="modelValue"
        :placeholder="props.placeholder"
      />
      <div
        v-if="props.icon"
        class="absolute inset-y-0 left-0 flex items-center px-3 text-gray-500"
      >
        <Icon :icon="props.icon" class="w-5 h-5" />
      </div>
    </div>
    <div
      v-if="!!errorMessage && isTouched"
      class="mt-1 text-xs text-red-500 dark:text-red-500"
    >
      {{ errorMessage }}
    </div>
  </div>
</template>