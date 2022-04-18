namespace HUIT2022.Pages
{
	public static class Setting
	{
		var url = GetEnvironmentVariable("API_URL");
		public static readonly string HostAddr = url;
	}
}
