<script setup lang="ts">
import Input from '../../global/input/Input.vue';
import Password from '../../global/input/password.vue';
import Button from '../../global/button/index.vue'
import { useRouter } from 'vue-router';
import { useForm } from 'vee-validate';
import * as yup from "yup";
import { useAuth } from '../../../store/auth';

const router = useRouter()
const authStore = useAuth()

const validationSchema = yup.object({
    email: yup.string().label("Email").required(),
    password: yup.string().label("Password").required(),
})

const { handleSubmit, values } = useForm({
    initialValues: {
        email: "",
        password: ""
    },
    validationSchema
})

const handleRegister = handleSubmit(async () => {
  const success = await authStore.register(values.email, values.password);

  if (success) {
    router.push({ name: 'Login' });
  } else {
    console.error(authStore.error || 'Register failed');
    alert(authStore.error || 'Register failed');
  }
});

const goToLoginPage = () => {
  router.push({ name: 'Login'})
}

</script>

<template>
    <section class="bg-gray-100 dark:bg-gray-900">
        <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
            <a href="#" class="flex items-center mb-6 text-2xl font-bold text-gray-900 dark:text-white">
                <a href="#" class="flex items-center text-2xl font-bold text-gray-900 dark:text-white">
                <!-- <img class="w-40 object-cover " src="/public/logo-eduparx.png"
                    alt="logo"> -->
            </a>
            </a>
            <div
                class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1
                        class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                        Registration Account    
                    </h1>
                    <form class="space-y-4 md:space-y-6" action="#">
                        <div>
                            <Input label="Email" name="email" placeholder="name@company.com" primary />
                        </div>
                        <div>
                            <Password label="Password" name="password" placeholder="*********" primary />
                        </div>
                        <Button text="Register" type="primary" customClass="w-full bg-blue-700" @click="handleRegister()"/>
                        <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                            Have an account yet? <a href="#"
                                class="font-medium text-primary-600 hover:underline dark:text-primary-500 cursor-pointer"
                                @click="goToLoginPage()">Sign in</a>
                        </p>
                    </form>
                </div>
            </div>
        </div>
    </section>
</template>
