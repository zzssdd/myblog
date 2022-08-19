import Vue from "vue"
import { 
    Button,
    Form,
    Icon,
    Input,
    FormItem,
    Message,
    Container,
    Header,
    Main,
    Footer,
    Aside,
    Menu,
    MenuItem,
    Row,
    Submenu,
    Col,
    MenuItemGroup,
    Table,
    TableColumn,
    Image,
    MessageBox,
    Pagination,
    Dialog,
    Switch,
    Upload,
    Card,
    Select,
    Option
} from "element-ui"

Vue.use(Button)
Vue.use(Form)
Vue.use(Icon)
Vue.use(Input)
Vue.use(FormItem)
Vue.use(Container)
Vue.use(Header)
Vue.use(Main)
Vue.use(Footer)
Vue.use(Aside)
Vue.use(Menu)
Vue.use(MenuItem)
Vue.use(Row)
Vue.use(Submenu)
Vue.use(Col)
Vue.use(MenuItemGroup)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Image)
Vue.use(Pagination)
Vue.use(Dialog)
Vue.use(Switch)
Vue.use(Upload)
Vue.use(Card)
Vue.use(Select)
Vue.use(Option)

Vue.prototype.$confirm=MessageBox.confirm
Vue.prototype.$prompt = MessageBox.prompt
Vue.prototype.$message = Message
