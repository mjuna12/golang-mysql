<script setup lang="ts">
import { computed, reactive } from "vue";
import { useRouter, RouterLink } from "vue-router";
import { Icon } from "@iconify/vue";

type ButtonProps = {
  text?: string;
  type?: "none" | "primary" | "secondary" | "blue" | "error" | "success";
  size?: "sm" | "md" | "lg" | "xl" | "2xl";
  to?: string | Record<string, unknown>;
  href?: string;
  icon?: string;
  block?: boolean;
  disabled?: boolean;
  customClass?: string;
  loading?: boolean;
};

const props = defineProps<ButtonProps>();

const router = useRouter();

const defaultStyle =
  "cursor-pointer flex justify-center focus:ring-gray-50 border gap-2 items-center shrink-0 shadow focus:ring-4 disabled:opacity-75 disabled:cursor-not-allowed";

const styles = reactive<Record<string, string>>({
  none: "",
  primary:
    "text-white bg-primary-600 border-primary-600 hover:bg-primary-700 hover:border-primary-700 focus:bg-primary-700",
  secondary:
    "bg-white text-slate-700 border-gray-300 hover:bg-gray-50 focus:bg-white",
  blue: "bg-primary-50 text-primary-700 border-primary-500 hover:bg-primary-50 focus:bg-primary-100",
  error:
    "text-white bg-error-600 border-error-600 hover:bg-error-700 hover:border-error-700 focus:bg-error-700",
  success: "text-white bg-teal-500  hover:bg-teal-700 focus:bg-teal-700",
});

const sizes = reactive<Record<string, string>>({
  md: "text-sm px-3 py-2.5 rounded-md font-semibold",
  sm: "text-xs py-2 px-3 rounded-md font-semibold",
  lg: "px-5 py-3 rounded-lg text-md font-semibold",
  xl: "px-5 py-3 rounded-lg font-semibold text-md",
  "2xl": "rounded-lg px-7 py-4 my-4 text-lg font-semibold",
});

const selectedStyle = computed(() => styles[props.type || "none"]);
const selectedSize = computed(() => sizes[props.size || "md"]);

const onClick = (event: MouseEvent) => {
  if (props.to) {
    router.push(props.to);
  } else if (!props.href) {
    event.preventDefault();
  }
};
</script>

<template>
  <RouterLink v-if="to" :to="to" custom>
    <template #default="{ navigate }">
      <a
        :class="[
          defaultStyle,
          selectedStyle,
          selectedSize,
          block ? 'w-full' : '',
          disabled ? 'cursor-not-allowed opacity-75' : '',
          customClass,
        ]"
        @click="navigate"
      >
        <Icon v-if="icon" :icon="icon" class="w-6 h-6" />
        <slot>{{ text }}</slot>
      </a>
    </template>
  </RouterLink>
  <button
    v-else
    :href="href"
    @click="onClick"
    :disabled="disabled"
    :class="[
      defaultStyle,
      selectedStyle,
      selectedSize,
      block ? 'w-full' : '',
      disabled ? 'cursor-not-allowed opacity-75' : '',
      customClass,
    ]"
  >
    <Icon
      v-if="icon && !loading"
      :icon="icon"
      class="w-6 h-6"
      :class="{
        'w-6 h-6': size === 'md',
      }"
    />

    <Icon
      v-if="loading"
      icon="eos-icons:bubble-loading"
      class="w-6 h-6 animate-spin"
    />

    <slot>{{ text }}</slot>
  </button>
</template>